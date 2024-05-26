export default async function getAllAttendances(token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendances`, {
        method: "GET",
        mode: 'cors',
        headers: {
            authorization: `Bearer ${token}`,
        },
        next: {
            tags: ['attendances']
        }
    });
    if (!response.ok) {
        throw new Error("Cannot get attendances for some reason");
    }
    return await response.json();
}