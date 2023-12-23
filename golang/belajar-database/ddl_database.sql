create table users(id serial primary key, 
				   name varchar(255), 
				   email varchar(255), 
				   password varchar(255), 
				   role varchar(255) , 
				   created_at timestamp default current_timestamp, 
				   updated_at timestamp, 
				   is_deleted boolean default false);