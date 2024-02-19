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


    return(
        <>
        <Box style={{
            width: "100vw",
            height: "100vh",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
            }}>
            <Typography variant="h2" component={"div"}>회원가입</Typography>
            <TextField
            name="name"
            required
            autoFocus
            value={registerData.name}
            label="이름"
            onChange={handleChange}
            placeholder="정상훈"
            />
            <TextField 
            name="age"
            required
            value={registerData.age} 
            label="나이" 
            onChange={handleChange}
            placeholder="20"
            />
            <TextField 
            name="gender"
            required
            value={registerData.gender} 
            label="성별" 
            onChange={handleChange}
            placeholder="남성"
            />
            <Button onClick={handleRegister}>
                회원가입
            </Button>
            <Snackbar open={openSnackbar} autoHideDuration={6000} onClose={handleCloseSnackbar}>
                <Alert onClose={handleCloseSnackbar} severity={messageType} sx={{ width: '100%' }}>
                    {message}
                </Alert>
            </Snackbar>
        </Box>
        </>
    )

}

export default RegisterPage