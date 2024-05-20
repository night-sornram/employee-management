export default async function Checkin(token : string, employee_id : string){
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/attendance/check-in`,{
        method: "POST",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            employee_id :  employee_id,
            check_in : new Date().toLocaleString()
        })

    })

    if(!response.ok) {
        throw new Error("Failed to check in")
    }

    return await response.json();
}
