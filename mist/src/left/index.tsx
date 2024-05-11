import { listen } from "@tauri-apps/api/event";
import { useState, useEffect } from "react";

const Left = (): JSX.Element => {
    const [notificationCounter, setNotificationCounter] = useState<number>(0);

    useEffect(() => {
        listen<string>('notification', (event) => {
          console.log('Received event:', event.payload, notificationCounter, 'mm');
    
          setNotificationCounter(notificationCounter + 1);
        });
    
      }, []);

      
    return (
        <p>Notification Counter: {notificationCounter}</p>

    )
}

export default Left;
