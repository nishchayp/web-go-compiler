package main;

import "fmt";

func div(x, y int) (int, int) {
    // ReturnStmt
    a := x / y;
    b :=kk x % y;
    return a, b;
}

func main() {
    // Declaration
    var i, j int = 1, 2;
    k := 3;
    cpp, python3, golang := true, false, 3;
    var c, python,, java = true, false, 77;
    var (
        x, y int; 
        z float32;
        );
    
    var a int = 0;
    // SimpleStmt
    a++;

    // GoStmt
    // go a;

    // GotoStmt
    goto label1;

    // LabeledStmt
    label1: t := 3 ^ 4;

    // Block
    {
        b, c := 3 - 9, 43 / 3;
        
        // IfStmt
        if c > b {
            c--;
        } else if c == b {
            c = 4;
        } else {
            b++;
        };
    };

    // SwitchStmt
    v := 42;
    switch v {
        case 100:
            k = 100;
            // FallthroughStmt
            fallthrough;
        case 42:
            k = 42;
            // BreakStmt
            break;
        case 1:
            k = 1;
            fallthrough;
        default:
            k++;
    };

    // ForStmt
    sum := 0;
    for i := 0; i < 10; i++ {
        if (i == 5) {
            // ContinueStmt
            continue;
        };
        sum = sum + i;
    };

    // DeferStmt
    defer fmt.Println("bye world!");
    
    div(10, 3);
    fmt.Println(34E99);
    fmt.Println(..44E+12);
    fmt.Println(.44E+12i);
    fmt.Println(32i);
    fmt.Println('q');
    fmt.Println("hello world", i, j, python3, golang, c, java, x, y, cpp, python, z, t);
    // a += b;
};

