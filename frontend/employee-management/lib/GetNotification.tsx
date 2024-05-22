export default async function GetNotification(token: string, eid: string , role : string) {
    if (role === "admin") {
        const response = await fetch(`${process.env.NEXT_PUBLIC_NOTIFICATION_URL}/api/notifications`, {
            method: "GET",
            mode: "cors",
            headers: {
                authorization: `Bearer ${token}`,
            },
        })
        if (!response.ok) {
            throw new Error("Failed to get notifications")
        }
        
        return await response.json();
    }
    else {
        const response = await fetch(`${process.env.NEXT_PUBLIC_NOTIFICATION_URL}/api/notifications/employee/${eid}`, {
            method: "GET",
            mode: "cors",
            headers: {
                authorization: `Bearer ${token}`,
            },
        })

        if (!response.ok) {
                throw new Error("Failed to get notifications")
        }

        return await response.json();    }

   
}