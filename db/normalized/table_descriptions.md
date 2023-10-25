# Описание таблиц
## Workspace
Рабочее пространство, в котором хранятся доски.
    id                          W1
    name                        W2
    thumbnail_url               W3
    date_created                W4
    description                 W5

## Board
Доска Канбан, разделенная на столбцы.
    id                          B1
    id_workspace                W1
    name                        B2
    description                 B3
    date_created                B4
    thumbnail_url               B5

## Column
Столбец в доске, в котором хранятся задания можно менять их порядок.
    id                          Co1
    id_board                    B1
    name                        Co2
    description                 Co3
    list_position               Co4

## Role
Роль пользователя в рабочем пространстве.
    id                          R1
    name                        R2
    description                 R3

## Task
Задание, ему можно добавить сроки и менять их порядок в столбце.
    id                          Ta1
    id_column                   Co1
    name                        Ta2
    date_created                Ta3
    description                 Ta4
    start                       Ta5
    end                         Ta6
    list_position               Ta7

## Tag
Тэги, которые можно добавить к заданиям и настроить их цвет в формате hex.
    id                          T1
    name                        T2
    color                       T3

## User
Пользователь сервиса, для работы достаточно указать почту и пароль, остальные поля опциональны, пароль хранится в виде хэша.
    id                          U1
    email                       U2
    password_hash               U3
    name                        U4
    surname                     U5
    avatar_url                  U6
    description                 U7

## Task_template
Шаблоны для заданий, некоторые созданы разработчиками, пользователи могут создавать свои. Данные хранятся в формате json, т.к. это лучший способ хранить копии данных из таблицы Task.
    id                          TT1
    data                        TT2

## Board_template
Шаблоны для досок, некоторые созданы разработчиками, пользователи могут создавать свои. Данные хранятся в формате json, т.к. это лучший способ хранить копии данных из таблицы Board.
    id                          BT1
    data                        BT2

## Checklist
Список задач внутри задания, можно добавить несколько в одно задание и поменять их порядок.
    id                          C1
    id_task                     Ta1
    name                        C2
    list_position               C3

## Checklist_item
Одна мини-задача в чек-листе, которую можно пометить как завершенную.
    id                          CI1
    id_checklist                C1
    name                        CI2
    done                        CI3
    list_position               CI4

## User_Workspace
Связующая таблица отношения М2М между пользователем и рабочими пространствоми, также позволяет добавить пользователю его роли в них.
    id_user                     U1
    id_workspace                W1
    id_role                     R1

## Board_User
Связующая таблица отношения М2М между пользователем и досками, к которым у него есть доступ.
    id_user                     U1
    id_board                    B1

## Task_User
Связующая таблица отношения М2М между пользователем и порученным ему заданием.
    id_user                     U1
    id_task                     Ta1

## Task_embedding
Ссылка на файловое вложение в задание, желательно хранить не упоминая конкретного сайта, а только путь к файлу, чтобы было легче их переносить.
    id                          TE1
    id_user                     U1
    id_task                     Ta1
    url                         TE1

## Session
Текущая сессия пользователя, чтобы не нужно было вводить логин и пароль несколько раз.
    token                       S1
    id_user                     U1
    expiration_date             S2

## Tag_Task
Связующая таблица отношения М2М между тэгами и заданиями, помеченными ими.
    id_tag                      T1
    id_task                     Ta1

## Comment
Комментарий к заданию, к которому можно прикрепить файл и на который можно ответить.
    id                          Com1
    id_user                     U1
    id_task                     Ta1
    content                     Com2
    date_created                Com3

## Comment_Reply
Таблица с ссылками на ответы к комментариям, чтобы рекурсивно проходить по ним.
    id_reply                    CR
    id_comment                  Com1

## Favourite_boards
Таблицы в списке избранного для каждого пользователя.
    id_board                    B1
    id_user                     U1

## User_Task_template
Связующая таблица отношения М2М между пользователями и шаблонами заданий, к которым у них есть доступ.
    id_user                     U1
    id_template                 TT1

## User_Board_template
Связующая таблица отношения М2М между пользователями и шаблонами досок, к которым у них есть доступ.
    id_user                     U1
    id_template                 BT1

## Reaction
Реакции-эмодзи на комментарии.
    id                          Re1
    id_user                     U1
    id_comment                  Com1
    content                     Re2

## Comment_embedding
Ссылка на файловое вложение в комментарий, желательно хранить не упоминая конкретного сайта, а только путь к файлу, чтобы было легче их переносить.
    id                          CE1
    id_user                     U1
    id_comment                  Com1
    url                         CE2
