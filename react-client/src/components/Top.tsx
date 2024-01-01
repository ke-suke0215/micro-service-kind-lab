import React from 'react';
import Button from '@mui/material/Button';
import axios from 'axios';

const Top = () => {
  const handleClick = async () => {
    try {
      const response = await axios.get('http://localhost:8080/test');
      console.log(response.data); // または適切な処理を行う
    } catch (error) {
      console.error('Request failed', error);
    }
  };

  return (
    <div>
      <Button variant="contained" color="primary" onClick={handleClick}>
        Click Me!
      </Button>
    </div>
  );
}

export default Top;
