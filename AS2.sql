CREATE database if not exists orders_by;
use orders_by;

create table if not exists items(
	item_id int NOT NULL,
    item_code varchar(10) NOT NULL,
    description varchar(255) NOT NULL,
    quantity int NOT NULL,
    order_id int,
    
    constraint PK_ItemID PRIMARY KEY (item_id),
    Foreign key (order_id) REFERENCES orders(order_id)
);

create table if not exists orders(
	order_id int not null,
    customer_name varchar(255),
    ordered_at varchar(255),
    
    constraint PK_Order_id PRIMARY KEY (order_id)
);

insert into orders values
	(1, 'Erico','2021-09-15T23:35:11Z'),
    (2, 'Erica','2021-09-16T10:35:11Z'),
    (3, 'Felix','2021-09-16T12:35:11Z');
    
insert into items values
	(1, '123', 'IPhone 10X', 1,1),
    (2, '231', 'Xiaomi Redmi Note 10', 1, 1),
    (3, '231', 'Xiaomi Redmi Note 10', 1, 2),
    (4, '123', 'IPhone 10X', 1,3);

    

