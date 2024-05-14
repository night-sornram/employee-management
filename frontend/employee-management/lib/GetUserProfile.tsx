export default async function GetUserProfile(token:string){
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/me`,{
        method: "GET",
        headers: {
            authorization: `Bearer ${token}`,
        },

    })

    if(!response.ok) {
        throw new Error("Failed to fetch user profile")
    }
    const responseBody = await response.json();
    return responseBody;
}