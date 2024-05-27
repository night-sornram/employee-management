export default async function GetTodayCheckIn(eid : string, token : string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendances/check-today/${eid}`, {
        method: 'GET',
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get your check in for some reason");
    }

    return await response.json();
}

