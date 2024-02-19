// MapPage.tsx

import React, { useState, useEffect } from 'react';
import { getMap } from '../api/map';

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
    <div className="MapPage">
      <h2>Google Map</h2>
      {loading ? (
        <p>Loading...</p>
      ) : (
        mapURL && <iframe src={mapURL} width="600" height="450" title="Google Map"></iframe>
      )}
    </div>
  );
};

export default MapPage;
