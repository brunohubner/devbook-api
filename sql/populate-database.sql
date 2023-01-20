insert into
  users (name, nick, email, password)
values
  (
    "Bruno",
    "bruno",
    "bruno@mail.com",
    "$2a$10$UBzXe.VfGjK7mXjYpMS6x.P.LkX3Kwlp0qokn42kIp.8IRO8tHvoG"
  ),
  (
    "Ana",
    "ana",
    "ana@mail.com",
    "$2a$10$UBzXe.VfGjK7mXjYpMS6x.P.LkX3Kwlp0qokn42kIp.8IRO8tHvoG"
  ),
  (
    "Bia",
    "bia",
    "bia@mail.com",
    "$2a$10$UBzXe.VfGjK7mXjYpMS6x.P.LkX3Kwlp0qokn42kIp.8IRO8tHvoG"
  ),
  (
    "Leo",
    "leo",
    "leo@mail.com",
    "$2a$10$UBzXe.VfGjK7mXjYpMS6x.P.LkX3Kwlp0qokn42kIp.8IRO8tHvoG"
  ),
  (
    "Pedro",
    "pedro",
    "pedro@mail.com",
    "$2a$10$UBzXe.VfGjK7mXjYpMS6x.P.LkX3Kwlp0qokn42kIp.8IRO8tHvoG"
  );

insert into
  followers (user_id, follower_id)
values
  (1, 2),
  (1, 3),
  (2, 1),
  (2, 3),
  (2, 5),
  (3, 1),
  (3, 2),
  (3, 4),
  (4, 1),
  (4, 5),
  (5, 1),
  (5, 3);

insert into
  posts (title, content, author_id)
values
  (
    "Publicação do Bruno",
    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
    1
  ),
  (
    "Publicação do Ana",
    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
    2
  ),
  (
    "Publicação do Bia",
    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
    3
  ),
  (
    "Publicação do Leo",
    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
    4
  ),
  (
    "Publicação do Pedro",
    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
    5
  );