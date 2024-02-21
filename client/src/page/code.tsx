// MapPage.tsx

import React, { useState, useEffect } from 'react';
import { getCode } from '../api/code';
import { Box, Typography } from '@mui/material';
import Header from '../component/header';


const CodePage: React.FC = () => {
  const [codeData, setCodeData] = useState('');
  const reportID = "1";

  useEffect(() => {
    const getQRcode = async () => {
      try {

        const response = await getCode(reportID);
        
        setCodeData(response);
      } catch (error) {
        console.error('Error:', error);
      }
    };

    getQRcode();
  }, []);


  return (
    <>
    <Header />
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
        width: '100vw',
        textAlign: 'center',
        p: 3, // 패딩
      }}
    >
      <Typography variant="h4" gutterBottom component="div" sx={{
        fontWeight: 'bold', // Assuming the header should be bold as per the sketch
        marginBottom: '20px', // Adding space below the header
      }}>
        QR Code
      </Typography>
      <div>
      {codeData && <img src={`data:image/png;base64,${codeData}`} alt="QR Code" />}
      </div>
    </Box>
    </>
  );
};

export default CodePage;
