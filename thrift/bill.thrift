struct BillInfo {
    1: string billID,
    2: string billMonth,
    3: i64 amount,
    4: string userID,
}

service BillService {
    list<BillInfo> GetBillList(1:string userID),
}