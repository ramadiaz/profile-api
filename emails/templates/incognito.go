package templates

const IncognitoEmailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Secure Message</title>
    <style>
        body {
            font-family: 'Segoe UI', Arial, sans-serif;
            line-height: 1.6;
            color: #2d3748;
            background-color: #f7fafc;
            margin: 0;
            padding: 40px 20px;
        }
        .email-wrapper {
            max-width: 600px;
            margin: 0 auto;
            background: linear-gradient(135deg, #4834d4, #686de0);
            border-radius: 16px;
            padding: 3px;
            box-shadow: 0 10px 25px rgba(0,0,0,0.1);
        }
        .message-container {
            background: #ffffff;
            border-radius: 14px;
            overflow: hidden;
        }
        .header-banner {
            background: linear-gradient(135deg, #4834d4, #686de0);
            padding: 20px;
            color: #ffffff;
        }
        .header-title {
            font-size: 24px;
            font-weight: 600;
            margin: 0;
        }
        .message-content {
            padding: 30px;
            background-color: #ffffff;
            line-height: 1.8;
        }
        .message-footer {
            background: #f8fafc;
            padding: 20px;
            font-size: 13px;
            color: #718096;
            border-top: 1px solid #e2e8f0;
        }
        .security-badge {
            background: linear-gradient(135deg, #4834d4, #686de0);
            color: #ffffff;
            padding: 4px 12px;
            border-radius: 12px;
            font-size: 12px;
            font-weight: 500;
            display: inline-block;
        }
    </style>
</head>
<body>
    <div class="email-wrapper">
        <div class="message-container">
            <div class="header-banner">
                <h1 class="header-title">{{.Subject}}</h1>
            </div>
            
            <div class="message-content">
                {{.MessageBody}}
            </div>
            
            <div class="message-footer">
                <span class="security-badge">{{.SecurityLevel}}</span>
                <div>Sent: {{.SentDate}}</div>
                <div>Message ID: {{.MessageID}}</div>
            </div>
        </div>
    </div>
</body>
</html>
`
