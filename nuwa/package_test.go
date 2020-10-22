package nuwa

import (
	`encoding/json`
	`testing`
)

var pkgJsonData = `
{
    "type": "windows",
    "maxRetry": 5,
    "srcFile": {
        "storageType": "cos",
        "filename": "btq14vst1biac60hu1s0",
        "storage": {
            "preassignedURL": "https://yunke-test-1259785003.cos.ap-shanghai.myqcloud.com/%2Fproduct%2Frelease%2Fbtq14vst1biac60hu1s0?response-content-disposition=attachment%3B+filename%3Dbtq14vst1biac60hu1s0%3Bfilename%2A%3Dutf-8%27%27btq14vst1biac60hu1s0&sign=q-sign-algorithm%3Dsha1%26q-ak%3DAKIDEJ84h3zwShfL5b4C86oKaibDYcoj5kCP%26q-sign-time%3D1603268269%3B1634804269%26q-key-time%3D1603268269%3B1634804269%26q-header-list%3D%26q-url-param-list%3Dresponse-content-disposition%26q-signature%3D522a0b17698116620f2965297d87c8dc3796cb5e",
            "method": "GET"
        }
    },
    "destFile": {
        "storageType": "cos",
        "filename": "bu7utbcdtucj6740v05g",
        "storage": {
            "preassignedURL": "https://yunke-test-1259785003.cos.ap-shanghai.myqcloud.com/1295254859162849280%2Frelease%2Fbu7utbcdtucj6740v05g?sign=q-sign-algorithm%3Dsha1%26q-ak%3DAKIDEJ84h3zwShfL5b4C86oKaibDYcoj5kCP%26q-sign-time%3D1603268269%3B1634804269%26q-key-time%3D1603268269%3B1634804269%26q-header-list%3D%26q-url-param-list%3D%26q-signature%3D962869a03b03f0288fb23d98631bfc864492cc17",
            "method": "PUT"
        }
    },
    "replaces": [
        {
            "filename": "resources/lib/conf/conf.json",
            "type": "json",
            "value": {
                "elements": [
                    {
                        "path": "domain",
                        "value": "http://localhost:8080"
                    },
                    {
                        "path": "name.zh",
                        "value": "云视课堂"
                    }
                ]
            }
        }
    ],
    "notify": {
        "url": "http://localhost:8080/api/clients/1316641343275012096/packages/notifies",
        "scheme": "Bearer",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTYwMzg3MzA2OSwianRpIjoiYnU3dXRiY2R0dWNqNjc0MHYwNTAiLCJpYXQiOjE2MDMyNjgyNjksImlzcyI6ImxvY2FsaG9zdDo4MDgwIiwibmJmIjoxNjAzMjY4MjY5LCJzdWIiOiJ7XCJpZFwiOlwiMVwiLFwiY3JlYXRlZEF0XCI6XCIyMDIwLTEwLTIxIDE2OjE3OjQ5XCIsXCJ1cGRhdGVkQXRcIjpcIjIwMjAtMTAtMjEgMTY6MTc6NDlcIixcInVzZXJuYW1lXCI6XCJhZG1pbkB5dW5zaGljbGFzcy5jb21cIixcInBob25lXCI6XCIrODYtMTcwODk3OTI3ODRcIn0ifQ.KiaDw53C71swxEm2MzkmvNt0VQ9iPXp4Op6MGf8s_e8",
        "payload": "eyJjbGllbnRJZCI6IjEzMTY2NDEzNDMyNzUwMTIwOTYiLCJwYWNrYWdlVGltZXMiOjB9"
    },
    "packager": {
        "productName": "云视课堂",
        "productVersion": "0.3.2",
        "productPublisher": "云视课堂",
        "productWebSite": "http://localhost:8080",
        "runFileName": "云视课堂.exe",
        "shortcutName": "云视课堂",
        "installDirName": "",
        "installIcon": {
            "storageType": "cos",
            "filename": "/product/file/pc-startup-logo.ico",
            "storage": {
                "preassignedURL": "https://yunke-test-1259785003.cos.ap-shanghai.myqcloud.com/%2Fproduct%2Ffile%2Fpc-startup-logo.ico?response-content-disposition=attachment%3B+filename%3D%25E4%25BA%2591%25E8%25A7%2586%25E8%25AF%25BE%25E5%25A0%2582.ico%3Bfilename%2A%3Dutf-8%27%27%25E4%25BA%2591%25E8%25A7%2586%25E8%25AF%25BE%25E5%25A0%2582.ico&sign=q-sign-algorithm%3Dsha1%26q-ak%3DAKIDEJ84h3zwShfL5b4C86oKaibDYcoj5kCP%26q-sign-time%3D1603268269%3B1634804269%26q-key-time%3D1603268269%3B1634804269%26q-header-list%3D%26q-url-param-list%3Dresponse-content-disposition%26q-signature%3D4917dc05965bc316d55cfbf543ea44826afe7b55",
                "method": "GET"
            }
        },
        "uninstallIcon": {
            "storageType": "cos",
            "filename": "/product/file/pc-startup-logo.ico",
            "storage": {
                "preassignedURL": "https://yunke-test-1259785003.cos.ap-shanghai.myqcloud.com/%2Fproduct%2Ffile%2Fpc-startup-logo.ico?response-content-disposition=attachment%3B+filename%3D%25E4%25BA%2591%25E8%25A7%2586%25E8%25AF%25BE%25E5%25A0%2582.ico%3Bfilename%2A%3Dutf-8%27%27%25E4%25BA%2591%25E8%25A7%2586%25E8%25AF%25BE%25E5%25A0%2582.ico&sign=q-sign-algorithm%3Dsha1%26q-ak%3DAKIDEJ84h3zwShfL5b4C86oKaibDYcoj5kCP%26q-sign-time%3D1603268269%3B1634804269%26q-key-time%3D1603268269%3B1634804269%26q-header-list%3D%26q-url-param-list%3Dresponse-content-disposition%26q-signature%3D4917dc05965bc316d55cfbf543ea44826afe7b55",
                "method": "GET"
            }
        },
        "uninstallMessage": "",
        "uninstallFinishMessage": ""
    },
    "payload": ""
}
`

func TestPackageUnmarshal(t *testing.T) {
	req := &Package{}

	if err := json.Unmarshal([]byte(pkgJsonData), req); nil != err {
		t.Error(err)
	}
}
