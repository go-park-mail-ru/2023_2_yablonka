# Функциональные зависимости
## Таблица Tag:
    {T1} -> T1
    {T1} -> T2
    {T1} -> T3

## Связь Tag и Task:
    {T1, Ta1} -> T1
    {T1, Ta1} -> T2
    {T1, Ta1} -> T3
    {T1, Ta1} -> Ta1
    {T1, Ta1} -> Ta2
    {T1, Ta1} -> Ta3
    {T1, Ta1} -> Ta4
    {T1, Ta1} -> Ta5
    {T1, Ta1} -> Ta6
    {T1, Ta1} -> Ta7

## Таблица Checklist_Item:
    {CI1} -> CI1
    {CI1} -> CI2
    {CI1} -> CI3
    {CI1} -> CI4

## Принадлежание Checklist_Item к одному Checklist-у:
    {CI1} -> C1

## Таблица Checklist:
    {C1} -> C1
    {C1} -> C2
    {C1} -> C3

## Принадлежание Checklist к одному Task-у:
    {C1} -> Ta1

## Таблица Workspace:
    {W1} -> W1
    {W1} -> W2
    {W1} -> W3
    {W1} -> W4

## Таблица Board:
    {B1} -> B1
    {B1} -> B2
    {B1} -> B3
    {B1} -> B4
    {B1} -> B5

## Принадлежание Board к одному Workspace-у:
    {B1} -> W1

## Таблица Column:
    {Co1} -> Co1
    {Co1} -> Co2
    {Co1} -> Co3
    {Co1} -> Co4
    {Co1} -> Co5

## Принадлежание Column к одному Board-у:
    {Co1} -> B1

## Таблица Task:
    {Ta1} -> Ta1
    {Ta1} -> Ta2
    {Ta1} -> Ta3
    {Ta1} -> Ta4
    {Ta1} -> Ta5
    {Ta1} -> Ta6
    {Ta1} -> Ta7

## Принадлежание Task к одному Column-у:
    {Ta1} -> Co1

## Role User-a по его id и связанному Workspace-у:
    {U1, W1} -> R1

## User, которому доверен Task:
    {UR, Ta1} -> UR
    {UR, Ta1} -> UWA
    {UR, Ta1} -> U1
    {UR, Ta1} -> U2
    {UR, Ta1} -> U3
    {UR, Ta1} -> U4
    {UR, Ta1} -> U5
    {UR, Ta1} -> U6
    {UR, Ta1} -> U7
    {UR, Ta1} -> Ta1
    {UR, Ta1} -> Ta2
    {UR, Ta1} -> Ta3
    {UR, Ta1} -> Ta4
    {UR, Ta1} -> Ta5
    {UR, Ta1} -> Ta6
    {UR, Ta1} -> Ta7

## Таблица Role:
    {R1} -> R1
    {R1} -> R2
    {R1} -> R3

## Таблица Task_Embedding:
    {TE1} -> TE1
    {TE1} -> TE2

## Принадлежание Task_Embedding к одному Task-у:
    {TE1} -> Ta1

## Принадлежание Task_Embedding к одному User-у:
    {TE1} -> U1

## Таблица User:
    {U1} -> U1
    {U1} -> U2
    {U1} -> U3
    {U1} -> U4
    {U1} -> U5
    {U1} -> U6
    {U1} -> U7
    
## Любимые Board-ы User-а:
    {FB} -> U1
    {FB} -> U2
    {FB} -> U3
    {FB} -> U4
    {FB} -> U5
    {FB} -> U6
    {FB} -> U7
    {FB} -> FB
    {FB} -> B1
    {FB} -> B2
    {FB} -> B3
    {FB} -> B4
    {FB} -> B5

## Связь Board и User:
    {UWA, B1} -> UWA
    {UWA, B1} -> U1
    {UWA, B1} -> U2
    {UWA, B1} -> U3
    {UWA, B1} -> U4
    {UWA, B1} -> U5
    {UWA, B1} -> U6
    {UWA, B1} -> U7
    {UWA, B1} -> B1
    {UWA, B1} -> B2
    {UWA, B1} -> B3
    {UWA, B1} -> B4
    {UWA, B1} -> B5

## Таблица Task_Template:
    {TT1} -> TT1
    {TT1} -> TT2

## Таблица Board_Template:
    {BT1} -> BT1
    {BT1} -> BT2

## Таблица Session:
    {S1} -> S1
    {S1} -> S2

## Принадлежание Session к одному User-у:
    {S1} -> U1

## Таблица Reaction:
    {Re1} -> Re1
    {Re1} -> Re2

## Принадлежание Reaction к одному User-у:
    {Re1} -> U1

## Принадлежание Reaction к одному Comment-у:
    {Re1} -> Com1

## Таблица Comment:
    {Com1} -> Com1
    {Com1} -> Com2
    {Com1} -> Com3

## Принадлежание Comment к одному User-у:
    {Com1} -> U1

## Принадлежание Comment к одному Task-у:
    {Com1} -> Ta1

## Таблица Comment_Embedding:
    {CE1} -> CE1
    {CE1} -> CE2

## Принадлежание Comment_Embedding к одному User-у:
    {CE1} -> U1

## Принадлежание Comment_Embedding к одному Comment-у:
    {CE1} -> Com1

## Принадлежание Comment_Reply к одному Comment-у:
    {CR} -> Com1

# Условно неизбыточное покрытие
        {T1} -> T2 T3
        {T1, Ta1} -> T2 T3 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4
        {W1} -> W2 W3 W4
        {B1} -> B2 B3 B4 B5 W1 W2 W3 W4
        {Co1} -> Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4
        {Ta1} -> Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4
        {C1} -> C2 C3 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4
        {CI1} -> CI2 CI3 CI4 C1 C2 C3 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4
        {U1, W1} -> U2 U3 U4 U5 U6 U7 W2 W3 W4 R1 R2 R3
        {UWA, B1} -> U1 U2 U3 U4 U5 U6 U7 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R3
        {UR, Ta1} -> UWA U1, U2 U3 U4 U5 U6 U7 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R3
        {R1} -> R2 R3
        {TE1} -> TE2 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4 U1 U2 U3 U4 U5 U6 U7 R1 R2 R3
        {U1} -> U2 U3 U4 U5 U6 U7
        {FB} -> U2 U3 U4 U5 U6 U7 B1 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R3
        {TT1} -> TT2
        {BT1} -> BT2
        {S1} -> S2 U1 U2 U3 U4 U5 U6 U7
        {Re1} -> Re2 U1 U2 U3 U4 U5 U6 U7 Com1 Com2 Com3 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R3
        {Com1} -> Com2 Com3 U1 U2 U3 U4 U5 U6 U7 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R3
        {CR} -> Com1 Com2 Com3 U1 U2 U3 U4 U5 U6 U7 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R3
        {CE1} -> CE2 U1 U2 U3 U4 U5 U6 U7 Com1 Com2 Com3 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 Co1 Co2 Co3 Co5 B1 B2 B3 B4 B5 W1 W2 W3 W4 R1 R2 R31
        {T1 T2 T3 CI1 CI2 CI3 CI4 C1 C2 C3 W1 W2 W3 W4 B1 B2 B3 B4 B5 Co1 Co2 Co3 Co5 Ta1 Ta2 Ta3 Ta4 Ta5 Ta6 Ta7 R1 R2 R3 TE1 TE2 U1 U2 U3 U4 U5 U6 U7 TT1 TT2 BT1 BT2 S1 S2 Re1 Re2 Com1 Com2 Com3 CE1 CE2 CR FB UWA} -> Nil
        