package lang

const (
    IS_BLOCKED = "¿Se bloquea?"
    CREATE = "Crear"
    CREATE_PROCESS = "Crear proceso"
    CREATE_PARTITION = "Crear partición"
    NAME = "Nombre"
    TIME = "Tiempo"
    SIZE = "Tamaño"
    PARTITION = "Partición"
    PROCESSES = "Procesos"
    PARTITIONS = "Particiones"
    PARTITION_NUMBER = "Número de partición"
    PARTITION_SIZE = "Tamaño de partición"
    MAIN_MENU = "Menú principal"
    MAKE_TICK = "Realizar iteración"
    WITHOUT_FILTERS = "Sin filtros"
    START_PROCESSOR = "Iniciar procesador"
    RESET_PROCESSOR = "Reiniciar procesador"

    READY = "Listo"
    RUNNING = "En ejecución"
    BLOCKED = "Bloqueado"
    FINISHED = "Finalizado"

    YES = "Sí"
    NO = "No"

    PROCESSOR_FINISHED = "El procesador ha termindo"

    // States
    INPUT_LIST = "Lista de entradas"
    READY_LIST = "Lista de listos"
    RUNNING_LIST = "Lista de en-ejecución"
    BLOCKED_LIST = "Lista de bloqueados"
    FINISHED_LIST = "Lista de finalizados"

    // Transitions
    READY_TO_RUNNING = "Listos a en-ejecución"
    RUNNING_TO_READY = "Tiempo expirado"
    RUNNING_TO_BLOCKED = "En-ejecución a bloqueado"
    BLOCKED_TO_READY = "E/S terminado"
    RUNNING_TO_FINISHED = "En-ejecución a terminado"


    ERROR_INVALID_NUMBER = "Error: número inválido."
    ERROR_EMPTY_TEXT = "Error: texto vacío."
    ERROR_COULDNT_ADD_PROCESS = "Error: no se pudo añadir el proceso."
    ERROR_COULDNT_ADD_PARTITION = "Error: no se pudo añadir la partición."
    ERROR_PARTITION_NOT_SELECTED = "Error: no se ha seleccionado una partición."
)
