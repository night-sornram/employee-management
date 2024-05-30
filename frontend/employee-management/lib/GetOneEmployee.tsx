export default async function GetOneEmployee(token: string, eid : string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/api/employees/${eid}`, {
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