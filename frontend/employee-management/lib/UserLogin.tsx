export default async function UserLogIn(userId : string, userPassword : string) {
    const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_URL}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
            id : userId,
            password : userPassword
            // id: "E12345",
            // password: "123456"
        }),
    });

    if(!response.ok){
        console.log(userId + " " + userPassword);
        throw new Error('Failed to login');
    }

    return await response.json();
}
