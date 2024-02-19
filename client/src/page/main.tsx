const MainPage: React.FC = () =>{

    return (
    <>
        <div style={{
                display: "flex",
                justifyContent: "center", // 수평 중앙 정렬
                alignItems: "center", // 수직 중앙 정렬
                height: "calc(100vh - 64px)", // 헤더의 높이를 뺀 나머지 높이
                width: "100%", // 너비 100%
                textAlign: "center", // 텍스트 중앙 정렬
                fontSize: "4rem", // 폰트 크기를 크게
        }}>
           메인 페이지
        </div>
    </>
    )
}

export default MainPage;