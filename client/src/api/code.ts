import ApiManager from ".";

export const getCode = async (reportID :string) => {
    try {
        const response = await ApiManager.get("/code", {
            params: { reportID },
            responseType: 'arraybuffer' // 'blob'도 가능합니다.
        });
        // `response.data`는 이제 arraybuffer 형식의 데이터를 포함하고 있습니다.
        // 이 데이터를 Base64 문자열로 변환해야 할 수 있습니다.
        const base64 = btoa(
            new Uint8Array(response.data)
                .reduce((data, byte) => data + String.fromCharCode(byte), '')
        );
        return base64;
    } catch (error) {
        throw error;
    }
}

export const scanCode = async (reportID :string) => {
    try {
        const response = await ApiManager.post("/code", {reportID});
        return response.data;
    } catch (error) {
        throw error;
    }
}