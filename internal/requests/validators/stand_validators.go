package validators

func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["repeat_password"] = append(errs["repeat_password"], "两次密码输入不一致")
	}

	return errs
}
