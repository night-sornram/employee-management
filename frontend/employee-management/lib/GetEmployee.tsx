export default async function GetEmployee(token: string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/employees`, {
        method: "GET",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error("Cannot get employee data");
    }
    return await response.json();
    
}