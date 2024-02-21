import React, { useEffect } from 'react';
import { BrowserMultiFormatReader } from '@zxing/library';
import { Box } from '@mui/material';
import { scanCode } from '../api/code';

const ScanCodePage: React.FC = () => {

  useEffect(() => {
    const codeReader = new BrowserMultiFormatReader();
    // 카메라 장치 ID를 지정하지 않으면, 기본 카메라를 사용합니다.
    // 카메라 장치 ID를 얻기 위해서는 codeReader.getVideoInputDevices()를 사용할 수 있습니다.
    const selectedDeviceId = undefined;

    // 스캔 시작
    codeReader.decodeFromVideoDevice(selectedDeviceId || null, 'video', (result, err) => {
        if (result) {
          // 스캔 성공: result 객체에서 QR 코드 데이터를 얻을 수 있습니다.
          alert(result.getText());
          (async () => {
            const response = await scanCode(result.getText());
            if (response) {
              alert("도움 이력을 성공적으로 등록하셨습니다!")
              window.location.href = "/profile"
            }
          })();
        }

        if (err && err.name !== 'NotFoundException') {
          // 스캔 중 발생한 에러 처리 (NotFoundException은 QR 코드가 발견되지 않았을 때 발생합니다.)
          console.error(err);
        }
    });

    // 컴포넌트 언마운트 시 스캔 중지
    return () => {
      codeReader.reset();
    };
  }, []);

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
        width: '95vw',
        textAlign: 'center',
        p: 3, // 패딩
        position: 'relative', // QR 프레임을 위한 포지셔닝
      }}
    >
      <div style={{ position: 'relative' }}>
        <video
          id="video"
          style={{ width: '100%' }}
        ></video>
        {/* QR 코드 프레임 추가 */}
        <div
          style={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            border: '5px solid #FF0000', // 빨간색 프레임
            width: '300px', // 프레임의 가로 크기
            height: '300px', // 프레임의 세로 크기
            zIndex: 10, // 비디오 위에 표시되도록 z-index 설정
          }}
        ></div>
      </div>
    </Box>
  );
};

export default ScanCodePage
