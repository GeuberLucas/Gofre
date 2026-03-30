export interface Portfolio {
  id: number; // uint
  user_id: number; // int
  asset_id: number; // uint
  deposit_date: Date;
  broker: string; // string
  amount: number; // float64
  is_done: boolean; // bool
  description: string; // string
}
