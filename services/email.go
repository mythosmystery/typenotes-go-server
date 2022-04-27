package services

import (
	"fmt"
	"os"
)

func SendResetEmail(to, token string) error {
	body := fmt.Sprintf(`
		<html>
			<body>
				<p>
					To reset your password, please click the link below.
				</p>
				<a href="%s/reset?token=%s">Reset Password</a>
			</body>
		</html>
	`, os.Getenv("APP_URL"), token)
	if _, err := SendMail(to, "Reset Password", body); err != nil {
		return err
	}
	return nil
}
