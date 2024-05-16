export default async function ChangePhone(token:string, phone:string, employee_id : string){
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/employees/${employee_id}`,{
        method: "PUT",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "phone" :  phone
        })

    })

    if(!response.ok) {
        throw new Error("Failed to change email")
    }

    return await response.json();
}