export default async function GetMyAttendances (eid: string, token: string) {
    const response = await fetch(`http://127.0.0.1:8081/attendance/me/${eid}`, {
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