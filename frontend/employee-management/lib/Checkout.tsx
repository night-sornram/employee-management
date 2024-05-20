export default async function Checkout(token : string, aid : Number){
    const response = await fetch(`${process.env.NEXT_PUBLIC_ATTENDANCE_URL}/api/attendance/check-out`,{
        method: "PUT",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "id" :  aid,
        })

    })

    if(!response.ok) {
        throw new Error("Failed to check in")
    }

    return await response.json();
}
