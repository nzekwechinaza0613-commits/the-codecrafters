  <!-- ═══════════════════════════════════════════
     SQUAD PIPELINE CONTRACT
        Squad: [BenchMarkers]
     ───────────────────────────────────────────
     Input line types:
       [line 1: ALL CAP]
       [line 2: lower]
       [line 3: Trimspace]
       [line 4: TODO ]
    
     Transformation rules (in order):
       1. [Trim all leading and trailing whitespace ]
       2. [Replace TODO: with ✦ ACTION]
       3. [Convert ALL CAPS lines to Title Case ]
       4. [Convert all lowercase lines to uppercase ]
       5. [Remove lines that are only dashes or blanks 	]
    
     Output format:
       [Header: yes/no — SENTINEL FIELD REPORT — PROCESSED]
       [Line numbering format- 1]
  
     Terminal summary fields:
        ✦ Lines read    : [number]                  						
       ✦ Lines written : [number]                  						
      ✦ Lines removed : [number]                  					
       ✦ Rules applied :   
       1. [Trim all leading and trailing whitespace ]
       2. [Replace TODO: with ✦ ACTION]
       3. [Convert ALL CAPS lines to Title Case ]
       4. [Convert all lowercase lines to uppercase ]
       5. [Remove lines that are only dashes or blanks 	]
     ═══════════════════════════════════════════ -->
