import React, { useState } from "react"
import { register } from "../api/oauth"
import { Box, Button, TextField, Typography, Snackbar } from "@mui/material"
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import axios from "axios"
import { ProfileParams } from "../model/profile";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });


const RegisterPage: React.FC = () =>{
    const [registerData, setRegisterData] = useState<ProfileParams>({
        name: "",
        age: 0,
        gender: "",
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
        const {name, value} = e.target
        if(name === "age"){
            setRegisterData({...registerData, [name]: parseInt(value)})
            return
        }
        setRegisterData({...registerData, [name]: value})
    }


    const [message, setMessage] = useState(""); // 메시지 상태
    const [messageType, setMessageType] = useState<"success" | "error">("success"); // 메시지 타입
    const [openSnackbar, setOpenSnackbar] = useState(false); // Snackbar의 열림 상태를 관리


    const handleRegister = async () => {
        
        
        try {
            await register(registerData);
            setMessage("회원가입이 성공적으로 완료되었습니다!");
            setMessageType("success");
            setOpenSnackbar(true);

            // 회원가입 성공 시 1초 후 로그인 페이지로 이동
            setTimeout(() => {
                window.location.href = "/oauth/google/login";
            }, 1000);
           
        } catch (error) {
            let errorMessage = "알 수 없는 에러가 발생했습니다";
            if (axios.isAxiosError(error)) {
                // Axios 오류 응답인 경우
                if (error.response && error.response.status === 400) {
                    // 서버로부터 받은 구체적인 오류 메시지 사용
                    errorMessage = error.response.data.error || "잘못된 요청입니다.";
                }
            } else if (error instanceof Error) {
                // 일반 JavaScript 오류인 경우
                errorMessage = error.message;
            }
            setMessage(errorMessage);
            setMessageType("error");
            setOpenSnackbar(true);
        }
    }

    const handleCloseSnackbar = (event: React.SyntheticEvent | Event, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setOpenSnackbar(false);
    };    


    return (
        <>
          <Box style={{
            width: "100vw",
            height: "100vh",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
            gap: "20px", // 입력 필드 간격 추가
            padding: "20px", // 패딩 추가
            boxSizing: "border-box", // 박스 크기 계산 방법 변경
          }}>
            <Typography variant="h2" component={"div"} style={{
              fontWeight: 'bold', // 글씨 굵기 변경
              marginBottom: '30px', // 제목 아래 여백 추가
            }}>
              회원가입
            </Typography>
            <TextField
              name="name"
              required
              autoFocus
              value={registerData.name}
              label="이름"
              onChange={handleChange}
              placeholder="정상훈"
              variant="outlined"
              style={{
                width: '20%', // 너비 설정
              }}
            />
            <TextField
              name="age"
              required
              value={registerData.age}
              label="나이"
              onChange={handleChange}
              placeholder="20"
              type="number"
              variant="outlined"
              style={{
                width: '20%', // 너비설정 
            }}/>
            <TextField
            name="gender"
            required
            value={registerData.gender}
            label="성별"
            onChange={handleChange}
            placeholder="남성"
            variant="outlined"
            style={{
                width: '20%', // 너비 설정
            }}
            />
            <Button
            onClick={handleRegister}
            variant="contained"
            style={{
            marginTop: '20px', // 버튼 위 여백 추가
            width: '15%', // 버튼 너비 설정
            padding: '10px 0', // 버튼 패딩 설정
            fontWeight: 'bold', // 글씨 굵기 변경
            }}
            >
            회원가입
            </Button>
            <Snackbar open={openSnackbar} autoHideDuration={6000} onClose={handleCloseSnackbar}>
            <Alert onClose={handleCloseSnackbar} severity={messageType} sx={{ width: '100%' }}>
            {message}
            </Alert>
            </Snackbar>
            </Box>
            </>
            );


}

export default RegisterPage