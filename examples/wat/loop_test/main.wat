(module
  (func (export "loop3") (result i32)
    (local $x i32)
    (block 
      (loop 
        get_local $x
        i32.const 1
        i32.add
        set_local $x
        (br_if 1 (i32.eq (get_local $x) (i32.const 3)))
        (br 0)
      )
    )
    get_local $x
  )
)