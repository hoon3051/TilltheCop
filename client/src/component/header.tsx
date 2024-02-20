import React, { useEffect, useState } from "react";
import { AppBar, Toolbar, IconButton, Typography, Box } from "@mui/material";
import { Link, useNavigate } from "react-router-dom";

// 세션에서 사용자 정보를 가져오는 함수
const checkTokenAndRedirect = () => {
  const userInfo = localStorage.getItem('accessToken');
  if (!userInfo) {
    return "/oauth/google/login";
  }
  try {
    // 세션 스토리지에서 문자열을 JSON 객체로 파싱
    return "/map";
  } catch (error) {
    console.error('Failed to parse user info from session storage:', error);
    return "/oauth/google/login";
  }
}

const Header: React.FC = () => {

  return (
    <AppBar component="nav" color="transparent">
      <Toolbar>
        <IconButton
          color="inherit"
          aria-label="open drawer"
          edge="start"
          sx={{ mr: 2, display: { sm: "none" } }}
        >
          {/* 아이콘 컴포넌트가 필요한 경우 여기에 추가 */}
        </IconButton>
        <Typography
            variant="h6"
            component="div"
            sx={{ flexGrow: 1, display: { xs: "none", sm: "block" } }}
          >
            <Link to={checkTokenAndRedirect()} style={{ color: 'black', textDecoration: 'none' }}>Till the Cop</Link>
          </Typography>
        <Box sx={{ display: { xs: "none", sm: "block" } }}>
          {/* 조건부 렌더링으로 사용자가 로그인한 경우에만 프로필 링크 표시 */}
        <Link to={"/profile"} style={{ color: 'black', textDecoration: 'none' }}>프로필</Link>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
