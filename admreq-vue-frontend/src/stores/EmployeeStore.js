import { defineStore } from "pinia";

export const useEmployeeStore = defineStore('EmployeeStore', {
        state: () => ({
                employeeTickets: [],
            }),
            actions: {
                async getEmployeeTickets(token) {
                    const res = await fetch("http://localhost:8080/v1/tsr/tickets",{
                        method: "POST",
                        body: JSON.stringify({
                            token: token,
                            mode: 'employee',
                        }),
                    })
                    const data = await res.json()
                    if (res.status === 200) {
                        console.log(data.tickets)
                        this.employeeTickets = data.tickets
                    }
                },
            }






            //async setEmployee(tsrid, emplid, token) {
        //    const res = await fetch("http://192.168.141.62:8080/v1/tsr/employee",{
        //        method: "POST",
        //        body: JSON.stringify({
        //            tsr_uuid: tsrid,
        //            employee_uuid: emplid,
        //            token: token
        //        })
        //    })
        //    if (res.status === 200) {
        //        this.getTickets(token)
        //    }
        //}

                //getUserTickets() {
        //    return this.tickets.filter((el) => el.userId === this.credentials.uuid)
        //},
        //getEmployeeTickets() {
        //    return this.tickets.filter((el) => el.employeeId === this.credentials.uuid)
        //},
        //setNoEmployeeTickets() {
        //    this.ticketsNotEmployee = this.tickets.filter((el) => el.employeeId === "")
        //},
})