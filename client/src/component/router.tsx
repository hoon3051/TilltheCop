import { Route, Routes } from "react-router-dom";
import MainPage from "../page/main";
import MapPage from "../page/map";
import LoginPage from "../page/login";
import CallbackPage from "../page/callback";



/**
 * 어느 url에 어떤 페이지를 보여줄지 정해주는 컴포넌트입니다.
 * Routes 안에 Route 컴포넌트를 넣어서 사용합니다.
 */

const RouteComponent = () => {
  return (
    <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/map" element={<MapPage />} />
        <Route path="/oauth/google/login" element={<LoginPage />} />
        <Route path="/oauth/google/callback" element={<CallbackPage />} />      
    </Routes>
  );
};

export default RouteComponent;