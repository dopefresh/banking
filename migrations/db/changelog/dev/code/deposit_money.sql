create or replace procedure banking.deposit_money(card_id int, summ decimal) as $$
begin
	update banking.card set balance = balance + $2 where card_id = $1;
	insert into banking.transaction(transaction_type, card_id, summ, transaction_datetime) VALUES('deposit', $1, $2, NOW());
end
$$ language plpgsql;
