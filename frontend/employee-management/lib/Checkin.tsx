export default async function Checkin(token : string, employee_id : string){
    console.log(employee_id)
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendance/check-in`,{
        method: "POST",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            eid : employee_id
        })

    })

    if(!response.ok) {
        throw new Error("Failed to check in")
    }

    return await response.json();
}
