export default async function ChangePassword(token:string, password : string, newPassword : string, id : string){
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/changePassword`,{
        method: "POST",
        mode : "cors",
        headers: {
            authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "id" : id,
            "password" : password,
            "new_password" : newPassword
        })

    })

    if(!response.ok) {
        throw new Error("Failed to change email")
    }

    return await response.json();
}