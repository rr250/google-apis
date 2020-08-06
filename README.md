# google-apis

1) Enable the Google APIs:- 
   * Visit https://developers.google.com/drive/api/v3/quickstart/go and click Enable Drive Api button and get credentials.json
   * Clone https://github.com/rr250/google-apis
   * Open google-apis repo
   * Paste your credentials.json here
   * Run ```go get```
   * Run ```go run quickstart.go```
   * It will prompt you to authorize access through a link which will appear in command-line promt
   * Browse to the provided URL in your web browser
   * If you are not already logged into your Google account, you will be prompted to log in
   * Accept all
   * Copy the code you're given, paste it into the command-line prompt, and press Enter
   * A token.json file will be created in the folder

2) I have put only sheets and drive apis. And some basic examples. To add others add api scope url ConfigFromJson line in main function of quickstart.go and delete token.json file and repeat 1st step.
