export default async function CreateAttendance (token: string, eid:string, dateTime: string, date: string, lid: number) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendances`,{
        method: "POST",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            employee_id : eid,
            check_in: dateTime,
            check_out: dateTime,
            date: date,
            leave_id: lid
        })

    })

    if(!response.ok) {
        throw new Error("Failed to create attendance")
    }

    return await response.json();
}