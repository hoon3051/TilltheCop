import ApiManager from ".";
import { ProfileParams } from "../model/profile";

export const getProfile = async () => {
    try {
        const response = await ApiManager.get("/profile");
        return response.data;
    } catch (error) {
        throw error;
    }
}

export const updateProfile = async (profile: ProfileParams) => {
    try {
        const response = await ApiManager.put("/profile", profile);
        return response.data;
    } catch (error) {
        throw error;
    }
}