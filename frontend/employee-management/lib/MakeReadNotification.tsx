export default async function MakeReadNotification(token: string, nid: Number) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_NOTIFICATION_URL}/api/notifications/read/${nid}`, {
        method: "PUT",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
            body: JSON.stringify({
            read: true,
        }),
    })
    return response.json()
    }