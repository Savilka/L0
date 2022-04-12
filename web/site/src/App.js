import './App.css';
import {useState} from "react";

function App() {
    const [id, setId] = useState("");
    const [order, setOrder] = useState(null);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        setId("");
        fetch("http://localhost:8080/order/?id=" + id)
            .then(response => {
                if (response.status === 200) response.json()
                    .then(data => setOrder(data))
                else setOrder(null)
            })
    };

    return (
        <div className="App">
            <form onSubmit={handleSubmit}>
                <label>
                    ID:
                    <input type="text" name="ID" value={id} onChange={(e => {
                        setId(e.target.value)
                    })} autoComplete={"off"}/>
                </label>
                <input type="submit" value="Search"/>
            </form>
            {order === null ? "empty" : <ul>
                <li>order_uid: {order.order_uid}</li>
                <li>track_number: {order.track_number}</li>
                <li>entry: {order.entry}</li>
                <li>delivery:
                    <ul>
                        <li>name: {order.delivery.name}</li>
                        <li>phone: {order.delivery.phone}</li>
                        <li>zip: {order.delivery.zip}</li>
                        <li>city: {order.delivery.city}</li>
                        <li>address: {order.delivery.address}</li>
                        <li>region: {order.delivery.region}</li>
                        <li>email: {order.delivery.email}</li>
                    </ul>
                </li>
                <li>payment:
                    <ul>
                        <li>transaction: {order.payment.transaction}</li>
                        <li>request_id: {order.payment.request_id}</li>
                        <li>currency: {order.payment.currency}</li>
                        <li>provide: {order.payment.provide}</li>
                        <li>amount: {order.payment.amount}</li>
                        <li>payment_dt: {order.payment.payment_dt}</li>
                        <li>bank: {order.payment.bank}</li>
                        <li>delivery_cost: {order.payment.delivery_cost}</li>
                        <li>goods_total: {order.payment.goods_total}</li>
                        <li>custom_fee: {order.payment.custom_fee}</li>
                    </ul>
                </li>
                {order.items.map((item) => {
                    return (
                        <li> items:
                            <ul>
                                <li>chrt_id: {item.chrt_id}</li>
                                <li>track_number: {item.track_number}</li>
                                <li>price: {item.price}</li>
                                <li>rid: {item.rid}</li>
                                <li>name: {item.name}</li>
                                <li>sale: {item.sale}</li>
                                <li>size: {item.size}</li>
                                <li>total_price: {item.total_price}</li>
                                <li>nm_id: {item.nm_id}</li>
                                <li>brand: {item.brand}</li>
                                <li>status: {item.status}</li>
                            </ul>
                        </li>

                    )
                })}
                <li>locale: {order.locale}</li>
                <li>internal_signature: {order.internal_signature}</li>
                <li>customer_id: {order.customer_id}</li>
                <li>delivery_service: {order.delivery_service}</li>
                <li>shardkey: {order.shardkey}</li>
                <li>sm_id: {order.sm_id}</li>
                <li>date_created: {order.date_created}</li>
                <li>oof_shard: {order.oof_shard}</li>
            </ul>}

        </div>
    );
}

export default App;
