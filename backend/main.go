/*********************************/
/*   LIDL_voice        ( //      */
/*   server main       ( )/      */
/*   by salade         )(/       */
/*  ________________  ( /)       */
/* ()__)____________)))))   :^}  */
/*********************************/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func startGeneration(app *pocketbase.PocketBase, source_name string, provided_text string) string {
	voice_path, err := app.Dao().FindFirstRecordByData("voices", "voice_name", source_name)
	if err != nil {
		return ""
	}
	arg1 := "output.wav"
	arg2 := provided_text
	arg3 := voice_path.GetString("voice_path")
	fmt.Printf("running run_python_tts.sh with args:\n[%s]\n[%s]\n[%s]\n", arg1, arg2, arg3)

	cmd := exec.Command("./run_python_tts.sh", arg1, arg2, arg3)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	fmt.Printf("Output: \n%s\n", output)
	return "output.wav"
}

// We use pocketbase as a framework to handle login
func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/createRecording",
			Handler: func(c echo.Context) error {
				source_filename := c.FormValue("voice")
				provided_text := c.FormValue("providedText")
				fmt.Printf("source_filename: %v, provided_text: %v\n", source_filename, provided_text)

				// startGeneration

				// if err != nil {
				//     return c.String(404, err.Error())
				// }
				startGeneration(app, source_filename, provided_text)
				wavBytes, err := os.ReadFile("output.wav")
				if err != nil {
					return c.String(http.StatusInternalServerError, "Failed to open generated file")
				}
				c.Response().Header().Set("Content-Type", "audio/wav")
				c.Response().Header().Set("Content-Dispoition", "attachment; filename=output.wav")

				return c.Blob(http.StatusOK, "audio/wav", wavBytes)
			},
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
