export interface UserJson {
    date_of_birth : string
    department : string
    email : string
    employee_id : string
    first_name_en : string
    first_name_th : string
    gender : string
    id : Number
    last_name_en : string
    last_name_th : string
    password : string
    phone : string
    role : string
    title_en : string
    title_th : string
}

export interface Attendance {
    id: number,
	employee_id: string,
	check_in: string,
	check_out: string,
	date: string,
	leave_id: number | null,
    employee_name: string,
    employee_lastname: string
}

export interface Leave {
    id: number,
	employee_id: string,
	date_start: string,
	date_end: string,
	reason: string,
	status: string,
    employee_name: string,
    employee_lastname: string
}

export interface Notification {
    id: number,
    employee_id: string,
    title: string,
    message: string,
    read: boolean,
}

export interface StateProp {
    notification : string,
    email : boolean
}

export interface FontProp{
    font : string
}