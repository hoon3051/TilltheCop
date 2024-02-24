import ApiManager from ".";
import { ProfileParams } from "../model/profile";

// Google OAuth 로그인 요청
export const googleLogin = async (callbackURL:string, registerURL:string) => {
    try {
        // 백엔드에 Google OAuth 로그인 요청을 보냅니다.
        const response = await ApiManager.get("/oauth/google/login?callbackURL=" + callbackURL + "&registerURL=" + registerURL);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export const register = async (profile: ProfileParams) => {
    try {
        // 백엔드에 회원가입 요청을 보냅니다.
        const response = await ApiManager.post("/oauth/register", profile);
        return response.data;
    } catch (error) {
        throw error;
    }
}
