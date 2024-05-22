export default async function CreateEmployee(token:string, data : any){
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/employees`,{
        method: "POST",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            ...data
        })

    })

    if(!response.ok) {
        throw new Error("Failed to change email")
    }

    return await response.json();
}