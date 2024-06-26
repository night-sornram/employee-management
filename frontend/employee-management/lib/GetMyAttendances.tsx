export default async function GetMyAttendances (eid: string, token: string,query: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendances/me/${eid}${query}`, {
        method: 'GET',
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get your attendances for some reason");
    }
    return await response.json();
}