export default async function GetMyAttendances (eid: string, token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/me/${eid}`, {
        method: 'GET',
        headers: {
            authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get your attendances for some reason");
    }
    return await response.json();
}