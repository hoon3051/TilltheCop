// MapPage.tsx

import React, { useState, useEffect } from 'react';
import { getProfile } from '../api/profile';
import { Box, Typography, Paper, List, ListItem, ListItemText, Button } from '@mui/material';
import { ProfileParams } from '../model/profile';
import Header from '../component/header';


const ProfilePage: React.FC = () => {
    const [profileData, setProfileData] = useState<ProfileParams>({
        name: "",
        age: 0,
        gender: "",
    })

  useEffect(() => {
    const getProfileData = async () => {
      try {
        const response = await getProfile()
        const { Name, Age, Gender } = response.profile; // 응답에서 프로필 데이터 추출

        setProfileData({ // 추출한 데이터로 상태를 업데이트합니다.
            name: Name,
            age: Age,
            gender: Gender,
          });
      } catch (error) {
        console.error('Error:', error);
      }
    };

    getProfileData();
  }, []);
  
    const goToUpdateProfile = () => {
        window.location.href = "/profile/update";
    }

  return (
    <>
    <Header/>
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
      <Typography variant="h2" gutterBottom component="div" sx={{
        fontWeight: 'bold', // Assuming the header should be bold as per the sketch
        marginBottom: '30px', // Adding space below the header
      }}>
        나의 프로필
      </Typography>
      <Paper elevation={3} sx={{ minWidth: '300px', p: 2 }}>
                <List>
                    <ListItem>
                        <ListItemText primary="이름" secondary={profileData.name} />
                    </ListItem>
                    <ListItem>
                        <ListItemText primary="나이" secondary={profileData.age} />
                    </ListItem>
                    <ListItem>
                        <ListItemText primary="성별" secondary={profileData.gender} />
                    </ListItem>
                </List>
            </Paper>
            <Button
            onClick={goToUpdateProfile}
            variant="contained"
            style={{
            marginTop: '20px', // 버튼 위 여백 추가
            width: '15%', // 버튼 너비 설정
            padding: '10px 0', // 버튼 패딩 설정
            fontWeight: 'bold', // 글씨 굵기 변경
            }}
            >
            프로필 수정
            </Button>
    </Box>
    </>
  );
};

export default ProfilePage;
