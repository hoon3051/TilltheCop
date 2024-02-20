// MapPage.tsx

import React, { useState, useEffect } from 'react';
import { getMap } from '../api/map';
import { Box, CircularProgress, Typography } from '@mui/material';
import Header from '../component/header';


const MapPage: React.FC = () => {
  const [mapURL, setMapURL] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const getLocationAndMap = async () => {
      try {
        setLoading(true);
        const { latitude, longitude } = await getCurrentLocation();
        const { map_url } = await getMap({ location_latitude: latitude, location_longitude: longitude });
        console.log('map_url:', map_url);
        setMapURL(map_url);
        setLoading(false);
      } catch (error) {
        console.error('Error:', error);
        setLoading(false);
      }
    };

    getLocationAndMap();
  }, []);

  const getCurrentLocation = (): Promise<{ latitude: number, longitude: number }> => {
    return new Promise((resolve, reject) => {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords;
          resolve({ latitude, longitude });
        },
        (error) => {
          reject(error);
        }
      );
    });
  };

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
        Google Map
      </Typography>
      {loading ? (
        <CircularProgress />
      ) : (
        mapURL && <iframe src={mapURL} width="60%" height="500px" style={{ border: 'none', boxShadow: '0 4px 8px 0 rgba(0,0,0,0.2)'  }} title="Google Map" />
      )}
    </Box>
    </>
  );
};

export default MapPage;
