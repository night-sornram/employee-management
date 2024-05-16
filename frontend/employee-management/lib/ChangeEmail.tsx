export default async function ChangeEmail(token:string, email:string, employee_id : string){
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/employees/${employee_id}`,{
        method: "PUT",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "email" :  email
        })

    })

    if(!response.ok) {
        throw new Error("Failed to change email")
    }

    return await response.json();
}