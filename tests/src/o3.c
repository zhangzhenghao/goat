float indirect_jump(const float *lhs, const float *rhs, long size)
{
    float result = 0;
    switch (size) {
    case 7:
        result += lhs[6] * rhs[6];
    case 6:
        result += lhs[5] * rhs[5];
    case 5:
        result += lhs[4] * rhs[4];
    case 4:
        result += lhs[3] * rhs[3];
    case 3:
        result += lhs[2] * rhs[2];
    case 2:
        result += lhs[1] * rhs[1];
    case 1:
        result += lhs[0] * rhs[0];
    }
    return result;
}
