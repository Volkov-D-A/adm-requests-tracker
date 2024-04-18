import { defineStore } from "pinia";

export const useUserStore = defineStore('UserStore', {
    state: () => ({
        credentials: [],
        authorized: false,
        userError: '',
        tickets: [],
        ticketsNotEmployee: [],
    }),
    actions: {
        async getAuth(user, password) {
            const res = await fetch("http://192.168.141.62:8080/v1/user/auth",{
                method: "POST",
                body: JSON.stringify({
                    login:user, 
                    password:password
                })
            })
            console.log(res)            
            const data = await res.json()
            if (res.status === 200) { 
                this.credentials = data
                console.log(this.credentials)
                this.authorized = true
            } else {
                this.userError = data.message
            }
        },
        async getTickets(token) {
            const res = await fetch("http://192.168.141.62:8080/v1/tsr/"+token,{
                method: "GET",
            })
            const data = await res.json()
            if (res.status === 200) {
                this.tickets = data.tickets
                this.setNoEmployeeTickets()
            }
        },
        getUserTickets() {
            return this.tickets.filter((el) => el.userId === this.credentials.uuid)
        },
        getEmployeeTickets() {
            return this.tickets.filter((el) => el.employeeId === this.credentials.uuid)
        },
        setNoEmployeeTickets() {
            this.ticketsNotEmployee = this.tickets.filter((el) => el.employeeId === "")
        },
        async createTicket(text, token) {
            const res = await fetch("http://192.168.141.62:8080/v1/tsr/create",{
                method: "POST",
                body: JSON.stringify({
                    text: text,
                    token: token
                }),
            })
            if (res.status === 200) {
                this.getTickets(token)
            }
        },
        async setEmployee(tsrid, emplid, token) {
            const res = await fetch("http://192.168.141.62:8080/v1/tsr/employee",{
                method: "POST",
                body: JSON.stringify({
                    tsr_uuid: tsrid,
                    employee_uuid: emplid,
                    token: token
                })
            })
            if (res.status === 200) {
                this.getTickets(token)
            }
        }
    }
})