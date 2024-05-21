export default async function CreateNotification(token: string, title: string, message: string, read : boolean, eid: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_NOTIFICATION_URL}/api/notifications`, {
        method: "POST",
        mode: "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            employee_id : eid,
            title: title,
            message: message,
            read: read,
        }),
    })

    if (!response.ok) {
        throw new Error("Failed to create notification")
    }

    return await response.json();
}