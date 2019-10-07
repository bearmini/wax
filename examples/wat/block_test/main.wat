(module
  (func $inRange (param $low i32) (param $high i32) (param $value i32) (result i32)
    (i32.and
      (i32.ge_s (get_local $value) (get_local $low))
      (i32.lt_s (get_local $value) (get_local $high))
    )  
  )

  (func (param $x i32) (param $y i32) (result i32)
    (block (result i32)
      ;; ensure that both the x and y value are within range:
      (i32.and
        (call $inRange
          (i32.const 0)
          (i32.const 50)
          (get_local $x)
        )
        (call $inRange
          (i32.const 0)
          (i32.const 50)
          (get_local $y)
        )
      )
    )
  )
  
  (export "block_test" (func 1))
)